import { observer } from 'mobx-react';
import { PlayerTable } from 'core/schemas';
import { worldState } from '..';
import Caribou from 'assets/Caribou.png';
import { ActivePositionWrapper } from 'components/PositionWrapper';

const Players = observer(() => {
  const players = PlayerTable.getAll(worldState.tableState);

  return (
    <>
      {players.map((player) => {
        return (
          <ActivePositionWrapper entity={player.Id} position={player.Position} key={player.Id}>
            <img src={Caribou} />
          </ActivePositionWrapper>
        );
      })}
    </>
  );
});

export default Players;
