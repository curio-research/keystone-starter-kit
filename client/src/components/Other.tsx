import styled from '@emotion/styled';
import { Position } from 'core/schemas';

export const tileSideWidth = 70;

interface PositionWrapperProps {
  position: Position;
}

export const PositionWrapper = styled.div<PositionWrapperProps>`
  position: absolute;
  left: ${(props) => props.position.x * tileSideWidth}px;
  bottom: ${(props) => props.position.y * tileSideWidth}px;
  width: ${tileSideWidth}px;
  height: ${tileSideWidth}px;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
`;
