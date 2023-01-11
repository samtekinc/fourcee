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

import { FaAws } from "react-icons/fa";
import { SiMicrosoftazure } from "react-icons/si";
import { FcGoogle } from "react-icons/fc";

import {
  CloudPlatform,
  ModuleAssignmentStatus,
} from "../__generated__/graphql";

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
          <BsFillCircleFill color="#20fc03" /> Active
        </>
      ) : moduleAssignmentStatus === ModuleAssignmentStatus.Inactive ? (
        <>
          <BsFillCircleFill color="#fc0303" /> Inactive
        </>
      ) : (
        <>
          <BsQuestionCircle color="#FF0000" /> Unknown
        </>
      )}
    </>
  );
}
