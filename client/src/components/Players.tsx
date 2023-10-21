import { observer } from 'mobx-react';
import { PlayerTable, TileTable } from '../core/schemas';
import { stateStore } from '..';
import { tileSideWidth } from './TerrainTiles';

const Players = observer(() => {
  const players = PlayerTable.getAll(stateStore.tableState.get(PlayerTable.name())!);

  console.log(players);

  return (
    <div>
      {players.map((player) => {
        return (
          <div
            key={player.Id}
            style={{
              position: 'absolute',
              left: player.Position.x * tileSideWidth,
              top: player.Position.y * tileSideWidth,
              width: `${tileSideWidth}px`,
              height: `${tileSideWidth}px`,
            }}
          ></div>
        );
      })}
    </div>
  );
});

export default Players;
