import { ModuleGroup } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { Breadcrumb, Col, Container, Row } from "react-bootstrap";
import {
  renderCloudPlatform,
  renderModuleAssignmentStatus,
} from "../utils/rendering";
import { NewModuleVersionButton } from "./NewModuleVersionButton";

const MODULE_GROUP_QUERY = gql`
  query moduleGroup($moduleGroupID: ID!) {
    moduleGroup(moduleGroupID: $moduleGroupID) {
      id
      name
      cloudPlatform
      versions {
        id
        name
        remoteSource
        terraformVersion
      }
      modulePropagations {
        name
        id
        moduleVersion {
          id
          name
        }
        orgUnit {
          id
          name
        }
        orgDimension {
          id
          name
        }
      }
      moduleAssignments {
        id
        moduleVersion {
          id
          name
        }
        modulePropagation {
          id
          name
        }
        orgAccount {
          id
          name
        }
        status
      }
    }
  }
`;

type Response = {
  moduleGroup: ModuleGroup;
};

export const ModuleGroupPage = () => {
  const params = useParams();

  const moduleGroupID = params.moduleGroupID ? params.moduleGroupID : "";

  const { loading, error, data, refetch } = useQuery<Response>(
    MODULE_GROUP_QUERY,
    {
      variables: {
        moduleGroupID: moduleGroupID,
      },
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container
      style={{ paddingTop: "2rem", maxWidth: "calc(100vw - 20rem)" }}
      fluid
    >
      <Breadcrumb>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/" }}>
          Home
        </Breadcrumb.Item>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/module-groups" }}>
          Modules
        </Breadcrumb.Item>
        <Breadcrumb.Item active>
          {data?.moduleGroup.name} ({data?.moduleGroup.id})
        </Breadcrumb.Item>
      </Breadcrumb>
      <Row>
        <Col md={"auto"}>
          <h1>
            {renderCloudPlatform(data?.moduleGroup.cloudPlatform)}{" "}
            {data?.moduleGroup.name}
          </h1>
        </Col>
      </Row>
      <h2>
        Versions{" "}
        <NewModuleVersionButton
          moduleGroupID={moduleGroupID}
          onCompleted={refetch}
        />
      </h2>
      <Table hover>
        <thead>
          <tr>
            <th>Module Version</th>
            <th>Terraform Version</th>
            <th>Remote Source</th>
          </tr>
        </thead>
        <tbody>
          {data?.moduleGroup.versions.map((version) => {
            return (
              <tr>
                <td>
                  <NavLink
                    to={`/module-groups/${data?.moduleGroup.id}/versions/${version?.id}`}
                  >
                    {version?.name}
                  </NavLink>
                </td>
                <td>{version?.terraformVersion}</td>
                <td>{version?.remoteSource}</td>
              </tr>
            );
          })}
        </tbody>
      </Table>
      <Row>
        <Col md={"auto"}>
          <h2>Module Propagations</h2>
          <Table hover>
            <thead>
              <tr>
                <th>Name</th>
                <th>Module Version</th>
                <th>Org Unit</th>
              </tr>
            </thead>
            <tbody>
              {data?.moduleGroup.modulePropagations.map((propagation) => {
                return (
                  <tr>
                    <td>
                      <NavLink to={`/module-propagations/${propagation?.id}`}>
                        {propagation?.name}
                      </NavLink>
                    </td>
                    <td>
                      <NavLink
                        to={`/module-groups/${data?.moduleGroup.id}/versions/${propagation?.moduleVersion?.id}`}
                      >
                        {propagation?.moduleVersion?.name}
                      </NavLink>
                    </td>
                    <td>
                      <NavLink
                        to={`/org-structures/${propagation?.orgDimension?.id}`}
                      >
                        {propagation?.orgDimension?.name}
                      </NavLink>
                      {" / "}
                      <NavLink
                        to={`/org-structures/${propagation?.orgDimension?.id}/org-units/${propagation?.orgUnit?.id}`}
                      >
                        {propagation?.orgUnit?.name}
                      </NavLink>
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </Table>
        </Col>
        <Col md={"auto"}>
          <h2>Module Assignments</h2>
          <Table hover>
            <thead>
              <tr>
                <th>Assignment ID</th>
                <th>Account</th>
                <th>Module Version</th>
                <th>Status</th>
                <th>Propagated By</th>
              </tr>
            </thead>
            <tbody>
              {data?.moduleGroup.moduleAssignments.map((assignment) => {
                return (
                  <tr>
                    <td>
                      <NavLink to={`/module-assignments/${assignment?.id}`}>
                        {assignment?.id}
                      </NavLink>
                    </td>
                    <td>
                      <NavLink
                        to={`/org-accounts/${assignment?.orgAccount?.id}`}
                      >
                        {assignment?.orgAccount?.name}
                      </NavLink>
                    </td>
                    <td>
                      <NavLink
                        to={`/module-groups/${moduleGroupID}/versions/${assignment?.moduleVersion?.id}`}
                      >
                        {assignment?.moduleVersion?.name}
                      </NavLink>
                    </td>
                    <td>{renderModuleAssignmentStatus(assignment?.status)}</td>
                    <td>
                      {assignment?.modulePropagation ? (
                        <NavLink
                          to={`/module-propagations/${assignment?.modulePropagation.id}`}
                        >
                          {assignment?.modulePropagation?.name}
                        </NavLink>
                      ) : (
                        <div>Direct Assignment</div>
                      )}
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </Table>
        </Col>
      </Row>
    </Container>
  );
};
