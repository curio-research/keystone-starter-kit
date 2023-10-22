import { observer } from 'mobx-react';
import { AnimalTable } from '../core/schemas';
import { worldState } from '..';
import { PositionWrapper } from './TerrainTiles';
import Duck from 'assets/Duck.png';

const Animals = observer(() => {
  const animals = AnimalTable.getAll(worldState.tableState);

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
