import { observer } from 'mobx-react';
import { ResourceTable } from '../core/schemas';
import { stateStore } from '..';
import { PositionWrapper } from './TerrainTiles';
import Meat from 'assets/Meat.png';

const Resources = observer(() => {
  const resources = ResourceTable.getAll(stateStore.tableState);

  return (
    <>
      {resources.map((resource) => {
        return (
          <PositionWrapper position={resource.Position} key={resource.Id}>
            <img src={Meat} />
          </PositionWrapper>
        );
      })}
    </>
  );
});

export default Resources;
