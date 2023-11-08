import {Box, Text, createStandaloneToast} from '@chakra-ui/react';
import { observer } from 'mobx-react';
import TerrainTile from '../components/TerrainTiles';
import Animals from '../components/Animals';
import Players from 'components/Players';
import { useEffect } from 'react';
import {Fire, Move} from 'core/requests';
import { useNavigate } from 'react-router-dom';
import Projectiles from 'components/Projectiles';
import {uiState} from 'index';
import Resources from 'components/Resources';

import {getPlayerID} from "../core/utils";

export const toast = createStandaloneToast();

// hard coded playerID. See InitGame.go TODO use random # generator
const playerId = -100;

// game page
const Game = observer(() => {
  const navigate = useNavigate();

  const handleKeyPress = (event: KeyboardEvent) => {
    const playerId = getPlayerID()
    if (playerId === undefined) {
      return
    }

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
        const lastPressedDirection = uiState.lastMovedDirection; // TODO so you need to move before firing?
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