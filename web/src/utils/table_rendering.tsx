import React from "react";
import OverlayTrigger from "react-bootstrap/OverlayTrigger";
import Tooltip from "react-bootstrap/Tooltip";
import Spinner from "react-bootstrap/Spinner";
import TimeAgo from "react-timeago";

import {
  BsXOctagonFill,
  BsCheckCircle,
  BsQuestionCircle,
  BsDashCircle,
  BsXCircle,
} from "react-icons/bs";

export function renderTimeField(time: string): JSX.Element {
  const timeDate = new Date(Date.parse(time));
  return (
    <>
      {time ? (
        <OverlayTrigger placement="top" overlay={<Tooltip>{time}</Tooltip>}>
          <td>
            <TimeAgo date={timeDate} />
          </td>
        </OverlayTrigger>
      ) : (
        <td>
          <BsXOctagonFill color="#FF0000" />
          <span>N/A</span>
        </td>
      )}
    </>
  );
}
