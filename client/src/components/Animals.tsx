import { observer } from 'mobx-react';
import { AnimalTable } from '../core/schemas';
import { stateStore } from '..';
import { PositionWrapper } from './TerrainTiles';

const Animals = observer(() => {
  const animals = AnimalTable.getAll(stateStore.tableState);

  return (
    <div>
      {animals.map((animal) => {
        return (
          <PositionWrapper
            position={animal.Position}
            key={animal.Id}
            style={{
              background: 'orange',
            }}
          />
        );
      })}
    </div>
  );
});

export default Animals;
