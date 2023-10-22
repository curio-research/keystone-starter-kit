import { observer } from 'mobx-react';
import { PlayerTable } from 'core/schemas';
import { stateStore } from '..';
import { PositionWrapper } from './TerrainTiles';
import Wolf from 'assets/Wolf.png';

const Players = observer(() => {
  const players = PlayerTable.getAll(stateStore.tableState);

  return (
    <>
      {players.map((player) => {
        return (
          <PositionWrapper position={player.Position} key={player.Id}>
            <img src={Wolf} />
          </PositionWrapper>
        );
      })}
    </>
  );
});

export default Players;
