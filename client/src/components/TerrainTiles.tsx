import styled from '@emotion/styled';
import { observer } from 'mobx-react';
import { Position, TileTable } from 'core/schemas';
import { worldState } from '..';
import Marsh from 'assets/Marsh.png';
import Bush from 'assets/Bush.png';
import Grass from 'assets/Grass.png';

export const tileSideWidth = 70;

const TerrainTile = observer(() => {
  const tiles = TileTable.getAll(worldState.tableState);

  return (
    <div>
      {tiles.map((tile) => {
        return (
          <>
            <PositionWrapper position={tile.Position} key={tile.Id}>
              <img src={Marsh} />
            </PositionWrapper>

            {tile.Terrain === false ? (
              <PositionWrapper position={tile.Position} key={tile.Id}>
                <img src={Bush} />
              </PositionWrapper>
            ) : (
              <PositionWrapper position={tile.Position} key={tile.Id}>
                <img src={Grass} style={{ opacity: 0.3 }} />
              </PositionWrapper>
            )}
          </>
        );
      })}
    </div>
  );
});

interface PositionWrapperProps {
  position: Position;
}

export const PositionWrapper = styled.div<PositionWrapperProps>`
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
