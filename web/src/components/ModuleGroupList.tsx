import React, { useState } from "react";
import { Maybe, ModuleVersion, ModuleGroup } from "../__generated__/graphql";
import { NavLink, Outlet, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Container from "react-bootstrap/Container";
import {
  Accordion,
  Button,
  Card,
  Col,
  Collapse,
  ListGroup,
  Nav,
  Row,
} from "react-bootstrap";
import { NewModuleGroupButton } from "./NewModuleGroupButton";
import { NewModuleVersionButton } from "./NewModuleVersionButton";
import { renderCloudPlatform } from "../utils/rendering";

interface ModuleVersionCollapseProps {
  moduleGroupID: string | undefined;
  moduleVersions: Maybe<ModuleVersion>[];
}

const ModuleVersionCollapse: React.VFC<ModuleVersionCollapseProps> = (
  props: ModuleVersionCollapseProps
) => {
  const [open, setOpen] = useState(false);

  return (
    <>
      <Button
        onClick={() => setOpen(!open)}
        aria-expanded={open}
        aria-controls="versions-collapse-list"
      >
        Versions
      </Button>
      <Collapse in={open}>
        <div id="versions-collapse-list">
          {props.moduleVersions.map((moduleVersion) => {
            return (
              <NavLink
                to={`/module-groups/${props.moduleGroupID}/versions/${moduleVersion?.id}`}
              >
                {moduleVersion?.name}
              </NavLink>
            );
          })}
        </div>
      </Collapse>
    </>
  );
};

const MODULE_GROUPS_QUERY = gql`
  query moduleGroups {
    moduleGroups(limit: 100) {
      id
      cloudPlatform
      name
      versions {
        id
        remoteSource
        terraformVersion
        name
      }
    }
  }
`;

type Response = {
  moduleGroups: ModuleGroup[];
};

export const ModuleGroupsList = () => {
  const params = useParams();

  const moduleGroupID = params.moduleGroupID ? params.moduleGroupID : "";

  const { loading, error, data, refetch } = useQuery<Response>(
    MODULE_GROUPS_QUERY,
    {
      pollInterval: 30000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container fluid>
      <Row className="flex-xl-nowrap">
        <Col
          md={2}
          xl={2}
          className="d-flex flex-column"
          style={{
            position: "sticky",
            top: "3.5rem",
            backgroundColor: "#f7f7f7",
            zIndex: 1000,
            maxWidth: "20rem",
            height: "calc(100vh - 3.5rem)",
            borderRight: "1px solid #dee2e6",
            paddingTop: "1rem",
          }}
        >
          <h3>Module Groups</h3>
          <Nav
            as={ListGroup}
            style={{
              maxHeight: "calc(100vh - 6rem)",
              flexDirection: "column",
              height: "100%",
              display: "flex",
              overflow: "auto",
              flexWrap: "nowrap",
            }}
          >
            {data?.moduleGroups.map((moduleGroup) => {
              return (
                <div style={{ padding: "0.25rem" }}>
                  <Card style={{}}>
                    <Card.Body>
                      <NavLink
                        to={`/module-groups/${moduleGroup?.id}`}
                        style={({ isActive }) =>
                          isActive
                            ? {
                                color: "blue",
                                textDecoration: "none",
                              }
                            : {
                                color: "inherit",
                                textDecoration: "none",
                              }
                        }
                      >
                        <Card.Title style={{ fontSize: "small" }}>
                          {renderCloudPlatform(moduleGroup?.cloudPlatform)}{" "}
                          {moduleGroup?.name}
                        </Card.Title>
                      </NavLink>
                      <Card.Text style={{}}>
                        Versions
                        {moduleGroup?.versions.map((version) => {
                          return (
                            <NavLink
                              to={`/module-groups/${moduleGroup?.id}/versions/${version?.id}`}
                              style={({ isActive }) =>
                                isActive
                                  ? {
                                      color: "blue",
                                      textDecoration: "none",
                                      paddingBottom: "0.1rem",
                                    }
                                  : {
                                      color: "inherit",
                                      textDecoration: "none",
                                      paddingBottom: "0.1rem",
                                    }
                              }
                            >
                              <Card>
                                <Card.Body>
                                  <Card.Title style={{ fontSize: "small" }}>
                                    {version?.name}
                                  </Card.Title>
                                  <Card.Text style={{ fontSize: "xx-small" }}>
                                    <b>TF Version: </b>
                                    {version?.terraformVersion}
                                    <br />
                                    {version?.remoteSource}
                                  </Card.Text>
                                </Card.Body>
                              </Card>
                            </NavLink>
                          );
                        })}
                      </Card.Text>
                    </Card.Body>
                  </Card>
                </div>
              );
            })}
            <br />
            <NewModuleGroupButton onCompleted={refetch} />
            <br />
          </Nav>
        </Col>
        <Col md={"auto"}>
          <Outlet />
        </Col>
      </Row>
    </Container>
  );
};
