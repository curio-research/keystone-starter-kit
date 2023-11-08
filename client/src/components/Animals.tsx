import { observer } from 'mobx-react';
import { AnimalTable } from '../core/schemas';
import { worldState } from '..';
import Duck from 'assets/Duck.png';
import { ActivePositionWrapper } from 'components/PositionWrapper';

const Animals = observer(() => {
  const animals = AnimalTable.getAll(worldState.tableState);

  return (
    <>
      {animals.map((animal) => {
        return (
          <ActivePositionWrapper entity={animal.Id} position={animal.Position} key={`animal-position-${animal.Id}`}>
            <img src={Duck} style={{ padding: '10px' }} key={`animal-img-${animal.Id}`} />
          </ActivePositionWrapper>
        );
      })}
    </>
  );
});

export default Animals;
