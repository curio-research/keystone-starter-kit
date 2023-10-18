import React from "react";
import { TableAccessor } from "../core/tableAccessor";
import { useSelector } from "react-redux";
import { StoreState, TableType, WorldType } from "../store/store";
import { Table, Thead, Tbody, Tfoot, Tr, Th, Td, TableCaption, TableContainer } from "@chakra-ui/react";

interface TableProps<T extends WithID> {
  table: TableType<T>;
  accessor: TableAccessor<T>;
}

export default function TableDisplay<T extends WithID>(props: { accessor: TableAccessor<T> }) {
  const accessor = props.accessor;
  const slice = useSelector((state: StoreState) => state.tableState.get(accessor.name()));
  if (slice === undefined) {
    return null;
  }

  return (
    <>
      <DisplayTable table={slice} accessor={accessor} />
    </>
  );
}

function DisplayTable<T extends WithID>(props: TableProps<T>) {
  const { table, accessor } = props;

  const anyVal = accessor.getAny(table);
  if (anyVal === undefined) {
    return null;
  }

  const columnNames = new Array<string>();
  for (const field in anyVal) {
    columnNames.push(field);
  }

  const allEntities = accessor.allEntities(table);

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
        const obj = accessor.get(table, entity)! as any;
        return (
          <Tbody key={entity}>
            <Tr key={entity}>
              {columnNames.map((columnName, index) => {
                return (
                  <Td key={index}>
                    <>{JSON.stringify(obj[columnName])}</>
                  </Td>
                );
              })}
            </Tr>
          </Tbody>
        );
      })}
    </Table>
  );
}

export interface WithID {
  Id: number;
}
