import { observer } from 'mobx-react';
import { ProjectileTable } from '../core/schemas';
import { worldState } from '..';
import Fire from 'assets/Fire.png';
import { ActivePositionWrapper } from 'components/PositionWrapper';

const Projectiles = observer(() => {
  const projectile = ProjectileTable.getAll(worldState.tableState);

  return (
    <>
      {projectile.map((projectile) => {
        return (
          <ActivePositionWrapper entity={projectile.Id} position={projectile.Position} key={`projectile-position-${projectile.Id}`}>
            <img src={Fire} key={`fire-img-${projectile.Id}`} />
          </ActivePositionWrapper>
        );
      })}
    </>
  );
});

export default Projectiles;
