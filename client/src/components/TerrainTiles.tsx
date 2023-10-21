import { observer } from 'mobx-react';
import { TileTable } from '../core/schemas';
import { stateStore } from '..';

export const tileSideWidth = 70;

const TerrainTile = observer(() => {
  const tiles = TileTable.getAll(stateStore.tableState.get(TileTable.name())!);

  return (
    <div>
      {tiles.map((tile) => {
        return (
          <div
            key={tile.Id}
            style={{
              position: 'absolute',
              left: tile.Position.x * tileSideWidth,
              top: tile.Position.y * tileSideWidth,
              width: `${tileSideWidth}px`,
              height: `${tileSideWidth}px`,
              background: tile.Terrain ? 'green' : 'brown',
            }}
          ></div>
        );
      })}
    </div>
  );
});

export default TerrainTile;
