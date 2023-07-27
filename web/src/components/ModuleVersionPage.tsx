import { ModuleVersion } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { Breadcrumb, Card, Col, Container, Row } from "react-bootstrap";
import { renderCloudPlatform, renderRemoteSource } from "../utils/rendering";

const MODULE_VERSION_QUERY = gql`
  query moduleVersion($moduleVersionID: ID!) {
    moduleVersion(moduleVersionID: $moduleVersionID) {
      id
      name
      moduleGroup {
        id
        cloudPlatform
        name
      }
      remoteSource
      terraformVersion
      variables {
        name
        type
        description
        default
      }
      modulePropagations {
        name
        description
        id
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
        name
        description
        orgAccount {
          id
          name
          cloudPlatform
        }
      }
    }
  }
`;

type Response = {
  moduleVersion: ModuleVersion;
};

export const ModuleVersionPage = () => {
  const params = useParams();

  const moduleGroupID = params.moduleGroupID ? params.moduleGroupID : "";
  const moduleVersionID = params.moduleVersionID ? params.moduleVersionID : "";

  const { loading, error, data } = useQuery<Response>(MODULE_VERSION_QUERY, {
    variables: {
      moduleGroupID: moduleGroupID,
      moduleVersionID: moduleVersionID,
    },
  });

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
        <Breadcrumb.Item
          linkAs={NavLink}
          linkProps={{
            to: `/module-groups/${data?.moduleVersion.moduleGroup.id}`,
          }}
        >
          {data?.moduleVersion.moduleGroup.name} (
          {data?.moduleVersion.moduleGroup.id})
        </Breadcrumb.Item>
        <Breadcrumb.Item
          linkAs={NavLink}
          linkProps={{
            to: `/module-groups/${data?.moduleVersion.moduleGroup.id}`,
          }}
        >
          Versions
        </Breadcrumb.Item>
        <Breadcrumb.Item active>
          {data?.moduleVersion.name} ({data?.moduleVersion.id})
        </Breadcrumb.Item>
      </Breadcrumb>

      <Row>
        <Col md={"auto"}>
          <h1>
            {renderCloudPlatform(data?.moduleVersion.moduleGroup.cloudPlatform)}{" "}
            {data?.moduleVersion.moduleGroup.name} {data?.moduleVersion.name}
          </h1>
        </Col>
      </Row>

      <p>
        <b>Remote Source:</b>{" "}
        {renderRemoteSource(data?.moduleVersion.remoteSource)}
        <br />
        <b>Terraform Version:</b> {data?.moduleVersion.terraformVersion}
      </p>
      <h2>Variables</h2>
      <Container fluid style={{ maxHeight: "50vh", overflow: "auto" }}>
        <Table striped hover bordered responsive>
          <thead>
            <tr>
              <th>Name</th>
              <th>Type</th>
              <th>Default</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            {data?.moduleVersion.variables.map((variable) => {
              return (
                <tr>
                  <td>{variable?.name}</td>
                  <td>{variable?.type}</td>
                  <td>{variable?.default}</td>
                  <td>{variable?.description}</td>
                </tr>
              );
            })}
          </tbody>
        </Table>
      </Container>
      <br />
      <h2>Module Propagations</h2>
      <Row>
        {data?.moduleVersion.modulePropagations.map((propagation) => {
          return (
            <Col md={"auto"}>
              <Card>
                <Card.Body>
                  <NavLink to={`/module-propagations/${propagation?.id}`}>
                    <Card.Title style={{ fontSize: "medium" }}>
                      {propagation?.name}
                    </Card.Title>
                  </NavLink>
                  <Card.Text style={{ fontSize: "small" }}>
                    <b>Org Unit: </b>
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
                    <br />
                    <b>Description: </b>
                    {propagation?.description}
                  </Card.Text>
                </Card.Body>
              </Card>
            </Col>
          );
        })}
      </Row>
    </Container>
  );
};
