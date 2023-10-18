package integration

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/curio-research/keystone/game/integration/testutils"
	pb_base "github.com/curio-research/keystone/game/proto/output/pb.base"
	pb_dict "github.com/curio-research/keystone/game/proto/output/pb.dict"
	pb_test "github.com/curio-research/keystone/game/proto/output/pb.test"
	"github.com/curio-research/keystone/game/server"
	"github.com/curio-research/keystone/game/systems/middleware"
	server2 "github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testBookTitle1  = "Cat in a Hat"
	testBookAuthor1 = "Dr. Seuss"

	testBookTitle2  = "Fault in Our Stars"
	testBookAuthor2 = "John Greene"

	testBookTitle3  = "The Order of the Phoenix"
	testBookAuthor3 = "J.K. Rowling"
)

var BookTable = state.NewTableAccessor[Book]()

type Book struct {
	Title   string
	Author  string
	OwnerID int
	Id      int
}

var p *testutils.PortManager

func init() {
	p = testutils.NewPortManager()
}

func TestAddBook(t *testing.T) {
	e, ws, s, _ := startServer(t)
	defer tearDown(ws, s)

	w := e.World

	playerID := 7
	player2ID := 8

	specificBookEntity := 969
	err := sendWSMsg(ws, playerID, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_AddSpecific,
		Title:  testBookTitle1,
		Author: testBookAuthor1,
		Entity: int64(specificBookEntity),
	})
	require.Nil(t, err)

	err = sendWSMsg(ws, playerID, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Add,
		Title:  testBookTitle2,
		Author: testBookAuthor2,
	})
	require.Nil(t, err)

	err = sendWSMsg(ws, player2ID, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Add,
		Title:  testBookTitle3,
		Author: testBookAuthor3,
	})
	require.Nil(t, err)

	b1 := BookTable.Get(w, specificBookEntity)
	assert.Equal(t, testBookTitle1, b1.Title)
	assert.Equal(t, testBookAuthor1, b1.Author)
	assert.Equal(t, playerID, b1.OwnerID)
	assert.Equal(t, specificBookEntity, b1.Id)

	b2 := BookTable.Filter(w, Book{
		Title:  testBookTitle2,
		Author: testBookAuthor2,
	}, []string{"Title", "Author"})
	require.Len(t, b2, 1)
	assert.Equal(t, playerID, BookTable.Get(w, b2[0]).OwnerID)

	b3 := BookTable.Filter(w, Book{
		Title:  testBookTitle3,
		Author: testBookAuthor3,
	}, []string{"Title", "Author"})
	require.Len(t, b3, 1)
	assert.Equal(t, player2ID, BookTable.Get(w, b3[0]).OwnerID)
}

func tearDown(ws *websocket.Conn, server *http.Server) {
	ws.Close()
	server.Close()
	time.Sleep(time.Second)
}

// also testing atomic tx
func TestUpdate(t *testing.T) {
	e, ws, s, mockErrorHandler := startServer(t)
	defer tearDown(ws, s)

	w := e.World

	playerID := 7

	b1Entity := addBook(w, testBookTitle1, testBookAuthor1, playerID) // entity = 1
	b2Entity := addBook(w, testBookTitle2, testBookAuthor2, playerID) // entity = 2

	// error; first update missing entity => none of the books should be updated
	err := sendWSMsg(ws, playerID, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Update,
		Title:  testBookTitle1,
		Author: testBookAuthor3,
	}, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Update,
		Title:  testBookTitle3,
		Author: testBookAuthor2,
		Entity: int64(b2Entity),
	})
	require.Nil(t, err)

	require.Equal(t, mockErrorHandler.ErrorCount(), 1)
	assert.Equal(t, "no book to update with entity 0", mockErrorHandler.LastError())

	b1 := BookTable.Get(w, b1Entity)
	assert.Equal(t, testBookTitle1, b1.Title)
	assert.Equal(t, testBookAuthor1, b1.Author)

	b2 := BookTable.Get(w, b2Entity)
	assert.Equal(t, testBookTitle2, b2.Title)
	assert.Equal(t, testBookAuthor2, b2.Author)

	err = sendWSMsg(ws, playerID, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Update,
		Title:  testBookTitle1,
		Author: testBookAuthor3,
		Entity: int64(b1Entity),
	}, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Update,
		Title:  testBookTitle3,
		Author: testBookAuthor2,
		Entity: int64(b2Entity),
	})
	require.Nil(t, err)

	b1 = BookTable.Get(w, b1Entity)
	assert.Equal(t, testBookTitle1, b1.Title)
	assert.Equal(t, testBookAuthor3, b1.Author)

	b2 = BookTable.Get(w, b2Entity)
	assert.Equal(t, testBookTitle3, b2.Title)
	assert.Equal(t, testBookAuthor2, b2.Author)
}

// also testing sending a call to another system
func TestDeleteAndFilter(t *testing.T) {
	playerID := 7
	player2ID := 8

	testTable := []struct {
		name         string
		authorFilter string
		titleFilter  string
		playerID     int
		errorMsg     string

		remainingEntities []int
	}{
		{
			name:              "one author",
			authorFilter:      testBookAuthor1,
			playerID:          playerID,
			remainingEntities: []int{2, 4, 5},
		},
		{
			name:              "author and title",
			titleFilter:       testBookTitle1,
			authorFilter:      testBookAuthor2,
			playerID:          player2ID,
			remainingEntities: []int{1, 3, 4, 5},
		},
		{
			name:              "no matching queries - playerID",
			authorFilter:      testBookAuthor1,
			playerID:          player2ID,
			remainingEntities: []int{1, 2, 3, 4, 5},
		},
		{
			name:              "no matching entities - filters",
			titleFilter:       testBookTitle1,
			authorFilter:      testBookAuthor3,
			playerID:          playerID,
			remainingEntities: []int{1, 2, 3, 4, 5},
		},
		{
			name:              "no author or title - error",
			playerID:          playerID,
			remainingEntities: []int{1, 2, 3, 4, 5},
			errorMsg:          "author or title must be provided to remove a book",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			e, ws, s, errorHandler := startServer(t)
			defer tearDown(ws, s)

			w := e.World

			addBookSpecific(w, testBookTitle1, testBookAuthor1, playerID, 1)  // entity = 1
			addBookSpecific(w, testBookTitle1, testBookAuthor2, player2ID, 2) // entity = 2
			addBookSpecific(w, testBookTitle2, testBookAuthor1, playerID, 3)  // entity = 3
			addBookSpecific(w, testBookTitle2, testBookAuthor2, player2ID, 4) // entity = 4
			addBookSpecific(w, testBookTitle3, testBookAuthor3, playerID, 5)  // entity = 5

			err := sendWSMsg(ws, testCase.playerID, &pb_test.TestBookInfo{
				Op:     pb_test.Operation_Remove,
				Title:  testCase.titleFilter,
				Author: testCase.authorFilter,
			})
			require.Nil(t, err)

			m := make(map[int]interface{})
			for _, i := range testCase.remainingEntities {
				m[i] = nil
			}

			for i := 1; i <= 5; i++ {
				book := BookTable.Get(w, i)
				if _, ok := m[i]; ok {
					assert.NotEqual(t, "", book.Title)
				} else {
					assert.Equal(t, "", book.Title)
				}
			}

			if testCase.errorMsg == "" {
				assert.Equal(t, errorHandler.ErrorCount(), 0)
			} else {
				require.Equal(t, errorHandler.ErrorCount(), 1)
				assert.Equal(t, testCase.errorMsg, errorHandler.LastError())
			}
		})
	}
}

func addBook(w state.IWorld, title, author string, ownerID int) int {
	return BookTable.Add(w, Book{
		Title:   title,
		Author:  author,
		OwnerID: ownerID,
	})
}

func addBookSpecific(w state.IWorld, title, author string, ownerID, entity int) int {
	return BookTable.AddSpecific(w, entity, Book{
		Title:   title,
		Author:  author,
		OwnerID: ownerID,
	})
}

func sendWSMsg(ws *websocket.Conn, playerID int, bookInfos ...*pb_test.TestBookInfo) error {
	err := testutils.SendMessage(ws, pb_dict.CMD_pb_test_C2S_Test, &pb_test.C2S_Test{
		BookInfos:       bookInfos,
		IdentityPayload: CreateMockIdentityPayload(playerID),
	})
	if err != nil {
		return err
	}

	time.Sleep(time.Second)
	return nil
}

// creates mock identity payload for testing purposes
func CreateMockIdentityPayload(playerId int) *pb_base.IdentityPayload {
	// construct identity payload
	return &pb_base.IdentityPayload{
		PlayerId: int64(playerId),
		JwtToken: "",
	}
}

func startServer(t *testing.T) (*server2.EngineCtx, *websocket.Conn, *http.Server, *testutils.MockErrorHandler) {
	mode := "dev" // TODO create enums for this?
	port, wsPort := p.GetPort(), p.GetPort()

	s, e, err := server.StartMainServer(mode, wsPort, "", 1)
	require.Nil(t, err)

	mockErrorHandler := testutils.NewMockErrorHandler()
	e.SystemErrorHandler = mockErrorHandler

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: s,
	}

	go func() {
		fmt.Println("starting server")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Errorf("http server closed with unexpected error %v", err)
			return
		}
	}()

	e.GameTick.Schedule.AddTickSystem(1, TestBookSystem)
	e.GameTick.Schedule.AddTickSystem(1, TestRemoveBookSystem)

	e.World.AddTable(BookTable)

	ws, err := testutils.SetupWS(t, wsPort)
	require.Nil(t, err)

	return e, ws, server, mockErrorHandler
}

var TestBookSystem = server2.CreateSystemFromRequestHandler(func(ctx *server2.TransactionCtx[*pb_test.C2S_Test]) {
	if ctx.GameCtx.Mode != "dev" {
		return
	}

	req := ctx.Req
	w := ctx.W
	playerID := int(req.GetIdentityPayload().GetPlayerId())

	for _, bookInfo := range req.BookInfos {
		switch bookInfo.Op {
		case pb_test.Operation_Add:
			BookTable.Add(w, Book{
				Title:   bookInfo.Title,
				Author:  bookInfo.Author,
				OwnerID: playerID,
			})
		case pb_test.Operation_AddSpecific:
			BookTable.AddSpecific(w, int(bookInfo.Entity), Book{
				Title:   bookInfo.Title,
				Author:  bookInfo.Author,
				OwnerID: playerID,
			})
		case pb_test.Operation_Remove:
			server2.QueueTxFromInternal(w, ctx.GameCtx.GameTick.TickNumber+1, testRemoveRequest{
				Title:    bookInfo.Title,
				Author:   bookInfo.Author,
				PlayerID: playerID,
			}, "")
		case pb_test.Operation_Update:
			book := BookTable.Get(w, int(bookInfo.Entity))
			if book.Title == "" {
				ctx.EmitError(fmt.Sprintf("no book to update with entity %v", bookInfo.Entity), []int{playerID})
				return
			}

			book.Title = bookInfo.Title
			book.Author = bookInfo.Author
			BookTable.Set(w, int(bookInfo.Entity), book)
		}
	}
}, middleware.VerifyIdentity[*pb_test.C2S_Test]())

type testRemoveRequest struct {
	Author   string
	Title    string
	PlayerID int
}

var TestRemoveBookSystem = server2.CreateSystemFromRequestHandler(func(ctx *server2.TransactionCtx[testRemoveRequest]) {
	req := ctx.Req
	w := ctx.GameCtx.World

	bookFilter := Book{Author: req.Author, Title: req.Title, OwnerID: req.PlayerID}
	fieldNames := []string{"OwnerID"}
	if req.Author == "" && req.Title == "" {
		ctx.EmitError("author or title must be provided to remove a book", []int{req.PlayerID})
		return
	}

	if req.Author != "" {
		fieldNames = append(fieldNames, "Author")
	}
	if req.Title != "" {
		fieldNames = append(fieldNames, "Title")
	}

	bookEntities := BookTable.Filter(w, bookFilter, fieldNames)
	for _, e := range bookEntities {
		BookTable.RemoveEntity(w, e)
	}
})