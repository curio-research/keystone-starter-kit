import {Box, Text, createStandaloneToast, position} from '@chakra-ui/react';
import { observer } from 'mobx-react';
import TerrainTile from '../components/TerrainTiles';
import Animals from '../components/Animals';
import Players from 'components/Players';
import { useEffect } from 'react';
import {CreatePlayer, Fire, Move} from 'core/requests';
import { useNavigate } from 'react-router-dom';
import Projectiles from 'components/Projectiles';
import {uiState, worldState} from 'index';
import Resources from 'components/Resources';
import {ethers} from "ethers/lib.esm";

export const toast = createStandaloneToast();

// hard coded playerID. See InitGame.go
const playerId = -100;
// TODO we need to find a way to get a new player ID for each player

// game page
export const playerIDKey = "existingPlayerID";

export const privateKey = "privateKey"

const Game = observer(() => {
  const navigate = useNavigate();

  // TODO put this in useEffect
  const existingPlayerID = localStorage.getItem(playerIDKey);
  if (existingPlayerID === "") {
    // TODO Do we assume that this is the authentication they want to use?
    const playerWallet = ethers.Wallet.createRandom();
    const newPlayerID = playerId; // TODO get this differently, either random or have endpoint that returns a playerID

    // TODO benefit of awaiting response is that we can make sure we only handle key press after player is set AND we can get a playerID
    CreatePlayer({PublicKey: playerWallet.publicKey, PlayerId: playerId})

    localStorage.setItem(playerIDKey, newPlayerID.toString());
    localStorage.setItem(privateKey, playerWallet.privateKey);
    // hope that table update will come quickly enough... But I guess it's the same regardless
  }


  const handleKeyPress = (event: KeyboardEvent) => {
    const playerId = parseInt(localStorage.getItem(playerIDKey)!, 10);
    switch (event.key) {
      case 'a':
        uiState.lastMovedDirection = 'left';
        Move({ Direction: 'left', PlayerId: playerId });
        break;

      case 's':
        uiState.lastMovedDirection = 'down';
        Move({ Direction: 'down', PlayerId: playerId });
        break;

      case 'd':
        uiState.lastMovedDirection = 'right';
        Move({ Direction: 'right', PlayerId: playerId });
        break;

      case 'w':
        uiState.lastMovedDirection = 'up';
        Move({ Direction: 'up', PlayerId: playerId });
        break;

      case ' ':
        const lastPressedDirection = uiState.lastMovedDirection;
        Fire({ Direction: lastPressedDirection, PlayerId: playerId });
        break;

      default:
        break;
    }
  };

  useEffect(() => {
    window.addEventListener('keydown', handleKeyPress);

    toast.toast({
      title: 'Welcome to the game!',
      description: 'Use WASD to move and space to shoot.',
      status: 'info',
      duration: 10_000,
      isClosable: true,
    });

    return () => {
      window.removeEventListener('keydown', handleKeyPress);
    };
  }, []);

  return (
    <Box p="10">
      <Text
        fontSize="sm"
        onClick={() => {
          navigate('/explore');
        }}
        style={{ cursor: 'pointer' }}
      >
        state explorer →
      </Text>

      <div style={{ height: '12px' }} />

      <div
        style={{
          width: '700px',
          height: '700px',
          position: 'relative',
        }}
      >
        <TerrainTile />
        <Animals />
        <Players />
        <Projectiles />
        <Resources />
      </div>
    </Box>
  );
});

export default Game;