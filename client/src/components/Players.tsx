import { observer } from 'mobx-react';
import { PlayerTable } from 'core/schemas';
import { worldState } from '..';
import { PositionWrapper } from './TerrainTiles';
import Caribou from 'assets/Caribou.png';

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
