import { observer } from 'mobx-react';
import { AnimalTable } from '../core/schemas';
import { worldState } from '..';
import Duck from 'assets/Duck.png';
import { ActivePositionWrapper } from 'components/Other';

const Animals = observer(() => {
  const animals = AnimalTable.getAll(worldState.tableState);

  return (
    <div>
      {animals.map((animal) => {
        return (
          <ActivePositionWrapper entity={animal.Id} position={animal.Position} key={animal.Id}>
            <img src={Duck} style={{ padding: '10px' }} />
          </ActivePositionWrapper>
        );
      })}
    </div>
  );
});

export default Animals;
