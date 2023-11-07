import { TableAccessor } from '../core/tableAccessor';
import { Table, Tbody, Tr, Th, Td, Text } from '@chakra-ui/react';
import { ITable } from '../store/types';
import { worldState } from '..';
import { observer } from 'mobx-react';

interface TableProps<T extends WithID> {
  table: ITable<T>;
  accessor: TableAccessor<T>;
}

// <T extends WithID>(props: { accessor: TableAccessor<T> }) => {
export const TableDisplay = observer(<T extends WithID>(props: { accessor: TableAccessor<T> }) => {
  const { accessor } = props;

  const tableName = worldState.tableState.get(accessor.name());
  if (tableName === undefined) {
    return null;
  }

  return (
    <>
      <DisplayTable table={tableName} accessor={accessor} />
    </>
  );
});

export default TableDisplay;

// <T extends WithID>(props: TableProps<T>) => {
const DisplayTable = observer(<T extends WithID>(props: TableProps<T>) => {
  const { table, accessor } = props;

  const anyVal = accessor.getAny(worldState.tableState);
  if (anyVal === undefined) {
    return null;
  }

  const columnNames = new Array<string>();
  for (const field in anyVal) {
    columnNames.push(field);
  }

  const allEntities = accessor.allEntities(worldState.tableState);

  return (
    <Table>
      <Tbody>
        <Tr>
          {columnNames.map((value) => {
            return <Th key={value}>{value}</Th>;
          })}
        </Tr>
      </Tbody>

      {allEntities.map((entity) => {
        const obj = accessor.get(worldState.tableState, entity)! as any;
        return (
          <Tbody key={entity}>
            <Tr key={entity}>
              {columnNames.map((columnName, index) => {
                return (
                  <Td key={index} style={{ padding: '10px' }}>
                    <Text fontSize="sm">{JSON.stringify(obj[columnName])}</Text>
                  </Td>
                );
              })}
            </Tr>
          </Tbody>
        );
      })}
    </Table>
  );
});

export interface WithID {
  Id: number;
}
