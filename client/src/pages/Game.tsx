import { Box } from '@chakra-ui/react';
import { observer } from 'mobx-react';
import TerrainTile from '../components/TerrainTiles';
import Animals from '../components/Animals';

// game page
const Game = observer(() => {
  return (
    <Box p="10">
      <div>/state explorer</div>
      <a href={'/explore'} style={{ height: '20px' }} />

      <div style={{ width: '700px', height: '700px', border: '1px solid black', position: 'relative' }}>
        <TerrainTile />
        <Animals />
      </div>
    </Box>
  );
});

export default Game;
