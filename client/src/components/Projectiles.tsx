import { observer } from 'mobx-react';
import { ProjectileTable } from '../core/schemas';
import { worldState } from '..';
import Fire from 'assets/Fire.png';
import { ActivePositionWrapper, PositionWrapper } from 'components/Other';

const Projectiles = observer(() => {
  const projectile = ProjectileTable.getAll(worldState.tableState);

  return (
    <>
      {projectile.map((projectile) => {
        return (
          <ActivePositionWrapper position={projectile.Position} key={projectile.Id}>
            <img src={Fire} />
          </ActivePositionWrapper>
        );
      })}
    </>
  );
});

export default Projectiles;
