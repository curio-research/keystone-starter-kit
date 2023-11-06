import { observer } from 'mobx-react';
import { ResourceTable } from '../core/schemas';
import { worldState } from '..';
import Meat from 'assets/Meat.png';
import { PositionWrapper } from 'components/Other';

const Resources = observer(() => {
  const resources = ResourceTable.getAll(worldState.tableState);

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
