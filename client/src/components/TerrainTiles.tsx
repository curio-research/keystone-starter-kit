import { observer } from 'mobx-react';
import { Position, TileTable } from 'core/schemas';
import { stateStore } from '..';
import styled from '@emotion/styled';

export const tileSideWidth = 70;

const TerrainTile = observer(() => {
  const tiles = TileTable.getAll(stateStore.tableState);

  return (
    <div>
      {tiles.map((tile) => {
        return (
          <PositionWrapper
            position={tile.Position}
            key={tile.Id}
            style={{
              background: tile.Terrain ? 'green' : 'brown',
            }}
          />
        );
      })}
    </div>
  );
});

interface IPositionWrapper {
  position: Position;
}

export const PositionWrapper = styled.div<IPositionWrapper>`
  position: absolute;
  left: ${(props) => props.position.x * tileSideWidth}px;
  top: ${(props) => props.position.y * tileSideWidth}px;
  width: ${tileSideWidth}px;
  height: ${tileSideWidth}px;
`;

export default TerrainTile;
