import { useEffect } from 'react';
import { Box, Select } from '@chakra-ui/react';
// eslint-disable-next-line max-len
// import { addTableUpdateToPendingUpdates, addUpdate, applyAllPendingUpdates, GetStateResponse, setIsFetchingState, setSelectedTableDisplay, store, StoreState, TableOperationType, TableUpdate } from "../store/store";
import { Accessors } from '../core/schemas';
import { observer } from 'mobx-react';
import Table from './Table';
import { uiStore } from '..';

const TableExplorer = observer(() => {
  return (
    <Box m={10}>
      <Box w="200px" mb={10}>
        <Select
          value={uiStore.selectedTableToDisplay || ''}
          placeholder="Select table"
          onChange={(e) => {
            uiStore.setSelectedTableToDisplay(e.target.value);
          }}
        >
          {Accessors.map((accessor) => {
            return (
              <option value={accessor.name()} key={accessor.name()}>
                {accessor.name()}
              </option>
            );
          })}
        </Select>
      </Box>

      {Accessors.map((accessor, index) => {
        return uiStore.selectedTableToDisplay === accessor.name() && <Table key={index} accessor={accessor} />;
      })}
    </Box>
  );
});

export default TableExplorer;
