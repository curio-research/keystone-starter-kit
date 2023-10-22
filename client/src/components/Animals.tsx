import { observer } from 'mobx-react';
import { AnimalTable } from '../core/schemas';
import { stateStore } from '..';
import { PositionWrapper } from './TerrainTiles';
import Duck from 'assets/DuckMale_east.png';

const Animals = observer(() => {
  const animals = AnimalTable.getAll(stateStore.tableState);

  return (
    <div>
      {animals.map((animal) => {
        return (
          <PositionWrapper position={animal.Position} key={animal.Id}>
            <img src={Duck} />
          </PositionWrapper>
        );
      })}
    </div>
  );
});

export default Animals;
