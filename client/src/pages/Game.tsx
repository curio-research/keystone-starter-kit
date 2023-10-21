import { Box } from '@chakra-ui/react';
import { observer } from 'mobx-react';
import TerrainTile from '../components/TerrainTiles';
import Animals from '../components/Animals';
import Players from 'components/Players';
import { useEffect } from 'react';
import { Move } from 'core/requests';

// TODO: remove
const playerId = -100;

const handleKeyPress = (event: any) => {
  switch (event.key) {
    case 'a':
      Move({ Direction: 'left', PlayerId: playerId });
      break;

    case 's':
      Move({ Direction: 'down', PlayerId: playerId });
      break;

    case 'd':
      Move({ Direction: 'right', PlayerId: playerId });
      break;

    case 'w':
      Move({ Direction: 'up', PlayerId: playerId });
      break;

    default:
      break;
  }
};

// game page
const Game = observer(() => {
  useEffect(() => {
    // Add an event listener for keydown events
    window.addEventListener('keydown', handleKeyPress);

    // Cleanup the event listener when the component unmounts
    return () => {
      window.removeEventListener('keydown', handleKeyPress);
    };
  }, []);

  return (
    <Box p="10">
      <div>/state explorer</div>
      <a href={'/explore'} style={{ height: '20px' }} />

      <div
        style={{
          width: '700px',
          height: '700px',
          border: '1px solid black',
          position: 'relative',
        }}
      >
        <TerrainTile />
        <Animals />
        <Players />
      </div>
    </Box>
  );
});

export default Game;
