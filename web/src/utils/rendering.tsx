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
  BsDashCircleFill,
  BsFillCircleFill,
} from "react-icons/bs";

import {
  MdSyncProblem,
  MdSync,
  MdOutlineAutoDelete,
  MdAddCircle,
} from "react-icons/md";

import { FaAws } from "react-icons/fa";
import { SiMicrosoftazure } from "react-icons/si";
import { FcGoogle } from "react-icons/fc";

import {
  CloudPlatform,
  ModuleAssignmentStatus,
  TerraformDriftCheckStatus,
} from "../__generated__/graphql";

export function renderApplyDestroy(destroy: boolean): JSX.Element {
  return (
    <>
      {destroy ? (
        <>
          <MdOutlineAutoDelete color="#F44336" /> Destroy
        </>
      ) : (
        <>
          <MdAddCircle color="#4CB950" /> Apply
        </>
      )}
    </>
  );
}

export function renderRemoteSource(
  remoteSource: string | undefined
): JSX.Element {
  let inputUrl = new URL("https://" + remoteSource || "");
  let version = inputUrl.searchParams.get("ref") || "main";

  // If the remote source is a reference to a subdirectory, we need to inject the version into the path at that point
  if (inputUrl.pathname.includes("//")) {
    let modifiedPath = inputUrl.pathname;
    modifiedPath = modifiedPath?.replace("//", `/tree/${version}/`);
    modifiedPath = modifiedPath;
    inputUrl.pathname = modifiedPath;
  } else {
    // we can just append the version to the end of the path
    inputUrl.pathname = inputUrl.pathname + `/tree/${version}`;
  }

  return (
    <a href={inputUrl.href} target="_blank">
      {remoteSource}
    </a>
  );
}

export function renderSyncStatus(
  status: TerraformDriftCheckStatus | undefined
): JSX.Element {
  return (
    <>
      {status === TerraformDriftCheckStatus.InSync ? (
        <>
          <MdSync color="#20fc03" />
          In Sync
        </>
      ) : status === TerraformDriftCheckStatus.OutOfSync ? (
        <>
          <MdSyncProblem color="#fc0303" />
          Out of Sync
        </>
      ) : status === TerraformDriftCheckStatus.Pending ? (
        <>
          <OverlayTrigger placement="top" overlay={<Tooltip>Pending</Tooltip>}>
            <Spinner animation="border" variant="secondary" size="sm" />
          </OverlayTrigger>{" "}
          Pending
        </>
      ) : (
        <>
          <BsQuestionCircle color="#FF0000" />
          Unknown
        </>
      )}
    </>
  );
}

export function renderCloudPlatform(
  cloudPlatform: CloudPlatform | undefined
): JSX.Element {
  return (
    <>
      {cloudPlatform === CloudPlatform.Aws ? (
        <FaAws color="FF9900" />
      ) : cloudPlatform === CloudPlatform.Azure ? (
        <SiMicrosoftazure color="007FFF" />
      ) : cloudPlatform === CloudPlatform.Gcp ? (
        <FcGoogle />
      ) : (
        <BsQuestionCircle color="#FF0000" />
      )}
    </>
  );
}

export function renderModuleAssignmentStatus(
  moduleAssignmentStatus: ModuleAssignmentStatus | undefined
): JSX.Element {
  return (
    <>
      {moduleAssignmentStatus === ModuleAssignmentStatus.Active ? (
        <>
          <BsFillCircleFill color="#4CB950" /> Active
        </>
      ) : moduleAssignmentStatus === ModuleAssignmentStatus.Inactive ? (
        <>
          <BsFillCircleFill color="#F44336" /> Inactive
        </>
      ) : (
        <>
          <BsQuestionCircle color="#FF0000" /> Unknown
        </>
      )}
    </>
  );
}
