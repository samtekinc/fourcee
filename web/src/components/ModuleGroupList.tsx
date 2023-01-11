import React, { useState } from "react";
import { ModuleGroups } from "../__generated__/graphql";
import { NavLink, Outlet, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Container from "react-bootstrap/Container";
import { Accordion, Col, ListGroup, Nav, Row } from "react-bootstrap";
import { NewModuleGroupButton } from "./NewModuleGroupButton";
import { NewModuleVersionButton } from "./NewModuleVersionButton";

const MODULE_GROUPS_QUERY = gql`
  query moduleGroups {
    moduleGroups(limit: 100) {
      items {
        moduleGroupId
        name
        versions {
          items {
            moduleVersionId
            name
          }
        }
      }
    }
  }
`;

type Response = {
  moduleGroups: ModuleGroups;
};

export const ModuleGroupsList = () => {
  const params = useParams();

  const moduleGroupId = params.moduleGroupId ? params.moduleGroupId : "";

  const { loading, error, data } = useQuery<Response>(MODULE_GROUPS_QUERY, {
    pollInterval: 1000,
  });

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

            height: "calc(100vh - 3.5rem)",
            borderRight: "1px solid #dee2e6",
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
            {data?.moduleGroups.items.map((moduleGroup) => {
              return (
                <Accordion defaultActiveKey={moduleGroupId}>
                  <Accordion.Item eventKey={moduleGroup?.moduleGroupId ?? ""}>
                    <Accordion.Header>
                      <NavLink
                        to={`/module-groups/${moduleGroup?.moduleGroupId}`}
                        style={({ isActive, isPending }) =>
                          isActive ? { fontWeight: 500 } : {}
                        }
                      >
                        {moduleGroup?.name}
                      </NavLink>
                    </Accordion.Header>
                    <Accordion.Body>
                      Versions
                      <ListGroup>
                        {moduleGroup?.versions.items.map((version) => {
                          return (
                            <ListGroup.Item>
                              <NavLink
                                to={`/module-groups/${moduleGroup?.moduleGroupId}/versions/${version?.moduleVersionId}`}
                                style={({ isActive, isPending }) =>
                                  isActive ? { fontWeight: 500 } : {}
                                }
                              >
                                {version?.name}
                              </NavLink>
                            </ListGroup.Item>
                          );
                        })}
                        <br />
                        <NewModuleVersionButton
                          moduleGroupId={moduleGroup?.moduleGroupId}
                        />
                      </ListGroup>
                    </Accordion.Body>
                  </Accordion.Item>
                </Accordion>
              );
            })}
            <NewModuleGroupButton />
          </Nav>
        </Col>
        <Col md={"auto"}>
          <Outlet />
        </Col>
      </Row>
    </Container>
  );
};
