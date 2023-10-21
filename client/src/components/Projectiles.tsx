import { observer } from 'mobx-react';
import { ProjectileTable } from '../core/schemas';
import { stateStore } from '..';
import { PositionWrapper } from './TerrainTiles';

const Projectiles = observer(() => {
  const projectile = ProjectileTable.getAll(stateStore.tableState);

  return (
    <div>
      {projectile.map((projectile) => {
        return (
          <PositionWrapper
            position={projectile.Position}
            key={projectile.Id}
            style={{
              background: 'red',
            }}
          />
        );
      })}
    </div>
  );
});

export default Projectiles;
