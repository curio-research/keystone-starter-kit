import { Box } from '@chakra-ui/react';
import { observer } from 'mobx-react';
import TerrainTile from '../components/TerrainTiles';

// game page
const Game = observer(() => {
  return (
    <Box p="10">
      <div style={{ width: '700px', height: '700px', border: '1px solid black', position: 'relative' }}>
        <TerrainTile />
      </div>
    </Box>
  );
});

export default Game;
