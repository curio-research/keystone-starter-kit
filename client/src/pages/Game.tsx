import { Box, Button } from '@chakra-ui/react';
import { observer } from 'mobx-react';
import TerrainTile from '../components/TerrainTiles';
import Animals from '../components/Animals';
import Players from 'components/Players';
import { useEffect } from 'react';
import { Fire, Move } from 'core/requests';
import { useNavigate } from 'react-router-dom';
import Projectiles from 'components/Projectiles';
import { uiStore } from 'index';

// TODO: remove
const playerId = -100;

// game page
const Game = observer(() => {
  const navigate = useNavigate();

  const handleKeyPress = (event: KeyboardEvent) => {
    switch (event.key) {
      case 'a':
        uiStore.lastMovedDirection = 'left';
        Move({ Direction: 'left', PlayerId: playerId });
        break;

      case 's':
        uiStore.lastMovedDirection = 'down';
        Move({ Direction: 'down', PlayerId: playerId });
        break;

      case 'd':
        uiStore.lastMovedDirection = 'right';
        Move({ Direction: 'right', PlayerId: playerId });
        break;

      case 'w':
        uiStore.lastMovedDirection = 'up';
        Move({ Direction: 'up', PlayerId: playerId });
        break;

      case ' ':
        const lastPressedDirection = uiStore.lastMovedDirection;
        Fire({ Direction: lastPressedDirection, PlayerId: playerId });
        break;

      default:
        break;
    }
  };

  useEffect(() => {
    window.addEventListener('keydown', handleKeyPress);

    return () => {
      window.removeEventListener('keydown', handleKeyPress);
    };
  }, []);

  return (
    <Box p="10">
      <Button
        onClick={() => {
          navigate('/explore');
        }}
      >
        state explorer →
      </Button>
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
      </div>
    </Box>
  );
});

export default Game;
