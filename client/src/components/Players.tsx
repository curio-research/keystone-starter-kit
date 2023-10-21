import { observer } from 'mobx-react';
import { PlayerTable } from 'core/schemas';
import { stateStore } from '..';
import { PositionWrapper } from './TerrainTiles';

const Players = observer(() => {
  const players = PlayerTable.getAll(stateStore.tableState);

  return (
    <div>
      {players.map((player) => {
        return (
          <PositionWrapper
            position={player.Position}
            key={player.Id}
            style={{
              backgroundColor: 'blue',
            }}
          />
        );
      })}
    </div>
  );
});

export default Players;
