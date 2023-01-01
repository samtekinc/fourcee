import React from "react";
import OverlayTrigger from "react-bootstrap/OverlayTrigger";
import Tooltip from "react-bootstrap/Tooltip";
import Spinner from "react-bootstrap/Spinner";
import TimeAgo from "react-timeago";
import { FcOk, FcHighPriority } from "react-icons/fc";

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

export function renderStatus(state: string | undefined): JSX.Element {
  if (state === "PENDING") {
    return (
      <>
        <OverlayTrigger placement="top" overlay={<Tooltip>Pending</Tooltip>}>
          <Spinner animation="border" variant="secondary" size="sm" />
        </OverlayTrigger>{" "}
        Pending
      </>
    );
  } else if (state === "RUNNING") {
    return (
      <>
        <OverlayTrigger placement="top" overlay={<Tooltip>Running</Tooltip>}>
          <Spinner animation="border" variant="info" size="sm" />
        </OverlayTrigger>{" "}
        Running
      </>
    );
  } else if (state === "SUCCEEDED") {
    return (
      <>
        <OverlayTrigger placement="top" overlay={<Tooltip>Running</Tooltip>}>
          <FcOk />
        </OverlayTrigger>{" "}
        Succeeded
      </>
    );
  } else if (state === "FAILED") {
    return (
      <>
        <OverlayTrigger placement="top" overlay={<Tooltip>Running</Tooltip>}>
          <FcHighPriority />
        </OverlayTrigger>{" "}
        Failed
      </>
    );
  }
  return (
    <>
      <BsQuestionCircle color="FF0000" /> {state}
    </>
  );
}
