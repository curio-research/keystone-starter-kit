import { observer } from 'mobx-react';
import { PlayerTable } from 'core/schemas';
import { worldState } from '..';
import Caribou from 'assets/Caribou.png';
import { PositionWrapper } from 'components/Other';

const Players = observer(() => {
  const players = PlayerTable.getAll(worldState.tableState);

  return (
    <>
      {players.map((player) => {
        return (
          <PositionWrapper position={player.Position} key={player.Id}>
            <img src={Caribou} />
          </PositionWrapper>
        );
      })}
    </>
  );
});

export default Players;
