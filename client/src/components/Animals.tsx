import { observer } from 'mobx-react';
import { AnimalTable } from '../core/schemas';
import { worldState } from '..';
import Duck from 'assets/Duck.png';
import { PositionWrapper } from 'components/Other';

const Animals = observer(() => {
  const animals = AnimalTable.getAll(worldState.tableState);

  return (
    <div>
      {animals.map((animal) => {
        return (
          <PositionWrapper position={animal.Position} key={animal.Id}>
            <img src={Duck} style={{ padding: '10px' }} />
          </PositionWrapper>
        );
      })}
    </div>
  );
});

export default Animals;
