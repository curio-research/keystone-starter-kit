import { observer } from 'mobx-react';
import { TileTable } from 'core/schemas';
import { worldState } from '..';
import Marsh from 'assets/Marsh.png';
import Bush from 'assets/Bush.png';
import Grass from 'assets/Grass.png';
import { PositionWrapper } from 'components/PositionWrapper';

const TerrainTile = observer(() => {
  const tiles = TileTable.getAll(worldState.tableState);

  return (
    <div>
      {tiles.map((tile) => {
        return (
          <>
            <PositionWrapper position={tile.Position} key={`marsh-tile-${tile.Id}`}>
              <img src={Marsh} key={`marsh-tile-img-${tile.Id}`} />
            </PositionWrapper>

            {tile.Terrain === false ? (
              <PositionWrapper position={tile.Position} key={`bush-tile-${tile.Id}`}>
                <img src={Bush} key={`marsh-bush-img-${tile.Id}`} />
              </PositionWrapper>
            ) : (
              <PositionWrapper position={tile.Position} key={`grass-tile-${tile.Id}`}>
                <img src={Grass} key={`marsh-grass-img-${tile.Id}`} style={{ opacity: 0.3 }} />
              </PositionWrapper>
            )}
          </>
        );
      })}
    </div>
  );
});

export default TerrainTile;
