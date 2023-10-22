import styled from '@emotion/styled';
import { observer } from 'mobx-react';
import { Position, TileTable } from 'core/schemas';
import { stateStore } from '..';
import Dirt from 'assets/Dirt.png';
import Bush from 'assets/BerryBush.png';

export const tileSideWidth = 70;

const TerrainTile = observer(() => {
  const tiles = TileTable.getAll(stateStore.tableState);

  return (
    <div>
      {tiles.map((tile) => {
        return (
          <>
            <PositionWrapper position={tile.Position} key={tile.Id}>
              <img src={Dirt} />
            </PositionWrapper>

            {tile.Terrain === false && (
              <PositionWrapper position={tile.Position} key={tile.Id}>
                <img src={Bush} />
              </PositionWrapper>
            )}
          </>
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
  bottom: ${(props) => props.position.y * tileSideWidth}px;
  width: ${tileSideWidth}px;
  height: ${tileSideWidth}px;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
`;

export default TerrainTile;
