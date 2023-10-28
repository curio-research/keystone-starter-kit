import { Box, Select } from '@chakra-ui/react';
import { observer } from 'mobx-react';
import TableDisplay from './Table';
import { uiState } from '..';
import { AllTableAccessors } from 'core/schemas';

const TableExplorer = observer(() => {
  return (
    <Box m={10}>
      <Box w="200px" mb={10}>
        <Select
          value={uiState.selectedTableToDisplay || ''}
          placeholder="Select table"
          onChange={(e) => {
            uiState.setSelectedTableToDisplay(e.target.value);
          }}
        >
          {AllTableAccessors.map((accessor) => {
            return (
              <option value={accessor.name()} key={accessor.name()}>
                {accessor.name()}
              </option>
            );
          })}
        </Select>
      </Box>

      {AllTableAccessors.map((accessor, index) => {
        return uiState.selectedTableToDisplay === accessor.name() && <TableDisplay key={index} accessor={accessor} />;
      })}
    </Box>
  );
});

export default TableExplorer;
