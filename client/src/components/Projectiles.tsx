import { observer } from 'mobx-react';
import { ProjectileTable } from '../core/schemas';
import { stateStore } from '..';
import { PositionWrapper } from './TerrainTiles';
import Fire from 'assets/Fire.png';

const Projectiles = observer(() => {
  const projectile = ProjectileTable.getAll(stateStore.tableState);

  return (
    <>
      {projectile.map((projectile) => {
        return (
          <PositionWrapper position={projectile.Position} key={projectile.Id}>
            <img src={Fire} />
          </PositionWrapper>
        );
      })}
    </>
  );
});

export default Projectiles;
