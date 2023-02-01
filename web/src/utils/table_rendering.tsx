import OverlayTrigger from "react-bootstrap/OverlayTrigger";
import Tooltip from "react-bootstrap/Tooltip";
import Spinner from "react-bootstrap/Spinner";
import TimeAgo from "react-timeago";
import { FcOk, FcHighPriority } from "react-icons/fc";

import { BsXOctagonFill, BsQuestionCircle } from "react-icons/bs";

import { CloudPlatform } from "../__generated__/graphql";
import { renderCloudPlatform } from "./rendering";

export function renderCloudPlatformTableData(
  cloudPlatform: CloudPlatform | undefined
): JSX.Element {
  return (
    <>
      <td>{renderCloudPlatform(cloudPlatform)}</td>
    </>
  );
}

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
      <div
        style={{
          color: "#6C757D",
        }}
      >
        <OverlayTrigger placement="top" overlay={<Tooltip>Pending</Tooltip>}>
          <Spinner animation="border" variant="secondary" size="sm" />
        </OverlayTrigger>{" "}
        Pending
      </div>
    );
  } else if (state === "RUNNING") {
    return (
      <div style={{ color: "#0DCAF0" }}>
        <OverlayTrigger placement="top" overlay={<Tooltip>Running</Tooltip>}>
          <Spinner animation="border" variant="info" size="sm" />
        </OverlayTrigger>{" "}
        Running
      </div>
    );
  } else if (state === "SUCCEEDED") {
    return (
      <div style={{ color: "#4CB950" }}>
        <OverlayTrigger placement="top" overlay={<Tooltip>Running</Tooltip>}>
          <FcOk />
        </OverlayTrigger>{" "}
        Succeeded
      </div>
    );
  } else if (state === "FAILED") {
    return (
      <div style={{ color: "#F44336" }}>
        <OverlayTrigger placement="top" overlay={<Tooltip>Running</Tooltip>}>
          <FcHighPriority />
        </OverlayTrigger>{" "}
        Failed
      </div>
    );
  }
  return (
    <>
      <BsQuestionCircle color="FF0000" /> {state}
    </>
  );
}
