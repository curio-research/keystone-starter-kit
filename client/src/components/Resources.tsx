import { observer } from 'mobx-react';
import { ResourceTable } from '../core/schemas';
import { stateStore } from '..';
import { PositionWrapper } from './TerrainTiles';

const Resources = observer(() => {
  const resources = ResourceTable.getAll(stateStore.tableState);

  return (
    <div>
      {resources.map((resource) => {
        return (
          <PositionWrapper
            position={resource.Position}
            key={resource.Id}
            style={{
              background: 'gold',
            }}
          />
        );
      })}
    </div>
  );
});

export default Resources;
