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
            <PositionWrapper position={tile.Position}>
              <img src={Marsh} />
            </PositionWrapper>

            {tile.Terrain === false ? (
              <PositionWrapper position={tile.Position}>
                <img src={Bush} />
              </PositionWrapper>
            ) : (
              <PositionWrapper position={tile.Position}>
                <img src={Grass} style={{ opacity: 0.3 }} />
              </PositionWrapper>
            )}
          </>
        );
      })}
    </div>
  );
});

export default TerrainTile;
