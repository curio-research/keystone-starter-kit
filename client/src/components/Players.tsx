import { observer } from 'mobx-react';
import { PlayerTable } from 'core/schemas';
import { worldState } from '..';
import { PositionWrapper } from './TerrainTiles';
import Wolf from 'assets/Wolf.png';

const Players = observer(() => {
  const players = PlayerTable.getAll(worldState.tableState);

  return (
    <>
      {players.map((player) => {
        return (
          <PositionWrapper position={player.Position} key={player.Id}>
            <img src={Wolf} style={{}} />
          </PositionWrapper>
        );
      })}
    </>
  );
});

export default Players;
