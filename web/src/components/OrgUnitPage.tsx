import { OrgUnit } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Container from "react-bootstrap/Container";
import { Card, Col, ListGroup, Row, Table } from "react-bootstrap";
import { renderCloudPlatform } from "../utils/rendering";
import { NewOrgUnitMembershipButton } from "./NewOrgUnitMembershipButton";
import { DeleteOrgUnitMembershipButton } from "./DeleteOrgUnitMembershipButton";

const ORG_UNIT_QUERY = gql`
  query orgUnit($orgUnitID: ID!) {
    orgUnit(orgUnitID: $orgUnitID) {
      id
      name

      orgDimension {
        id
        name
      }

      upstreamOrgUnits {
        id
        name
        modulePropagations {
          id
          name
          description
          moduleGroup {
            id
            name
            cloudPlatform
          }
          moduleVersion {
            id
            name
          }
        }
      }

      orgAccounts {
        id
        name
        cloudPlatform
        cloudIdentifier
      }

      modulePropagations {
        id
        name
        description
        moduleGroup {
          id
          name
          cloudPlatform
        }
        moduleVersion {
          id
          name
        }
      }
    }
  }
`;

type Response = {
  orgUnit: OrgUnit;
};

export const OrgUnitPage = () => {
  const params = useParams();

  const orgUnitID = params.orgUnitID ? params.orgUnitID : "";
  const orgDimensionID = params.orgDimensionID ? params.orgDimensionID : "";

  const { loading, error, data, refetch } = useQuery<Response>(ORG_UNIT_QUERY, {
    variables: {
      orgUnitID: orgUnitID,
      orgDimensionID: orgDimensionID,
    },
    pollInterval: 10000,
  });

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container
      fluid
      style={{ paddingTop: "2rem", borderTop: "1px solid black" }}
    >
      <h2>Module Propagations</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Name</th>
            <th>Module Group</th>
            <th>Module Version</th>
            <th>Source</th>
            <th>Description</th>
          </tr>
        </thead>
        <tbody>
          {data?.orgUnit.modulePropagations.map((modulePropagation) => {
            return (
              <tr>
                <td>
                  <NavLink to={`/module-propagations/${modulePropagation?.id}`}>
                    {modulePropagation.name}
                  </NavLink>{" "}
                  <sup>
                    {renderCloudPlatform(
                      modulePropagation?.moduleGroup.cloudPlatform
                    )}
                  </sup>
                </td>
                <td>
                  <NavLink
                    to={`/module-groups/${modulePropagation?.moduleGroup.id}`}
                  >
                    {modulePropagation.moduleGroup.name}{" "}
                  </NavLink>
                </td>
                <td>
                  <NavLink
                    to={`/module-groups/${modulePropagation?.moduleGroup.id}/versions/${modulePropagation?.moduleVersion.id}`}
                  >
                    {modulePropagation.moduleVersion.name}
                  </NavLink>
                </td>
                <td>Local</td>
                <td>{modulePropagation.description}</td>
              </tr>
            );
          })}
          {data?.orgUnit.upstreamOrgUnits.map((upstreamOrgUnit) => {
            return (
              <>
                {upstreamOrgUnit?.modulePropagations.map(
                  (modulePropagation) => {
                    return (
                      <>
                        <tr>
                          <td>
                            <NavLink
                              to={`/module-propagations/${modulePropagation?.id}`}
                            >
                              {modulePropagation.name}
                            </NavLink>
                            <sup>
                              {renderCloudPlatform(
                                modulePropagation?.moduleGroup.cloudPlatform
                              )}
                            </sup>
                          </td>
                          <td>
                            <NavLink
                              to={`/module-groups/${modulePropagation?.moduleGroup.id}`}
                            >
                              {modulePropagation.moduleGroup.name}{" "}
                            </NavLink>
                          </td>
                          <td>
                            <NavLink
                              to={`/module-groups/${modulePropagation?.moduleGroup.id}/versions/${modulePropagation?.moduleVersion.id}`}
                            >
                              {modulePropagation.moduleVersion.name}
                            </NavLink>
                          </td>
                          <td>
                            <NavLink
                              to={`/org-structures/${orgDimensionID}/org-units/${upstreamOrgUnit.id}`}
                            >
                              {upstreamOrgUnit?.name} OU
                            </NavLink>
                          </td>
                          <td>{modulePropagation.description}</td>
                        </tr>
                      </>
                    );
                  }
                )}
              </>
            );
          })}
        </tbody>
      </Table>

      <br />
      <h2>Org Account Membership</h2>
      <Table striped bordered hover>
        <thead>
          <th>Name</th>
          <th>Cloud Identifier</th>
          <th>Remove</th>
        </thead>
        <tbody>
          {data?.orgUnit.orgAccounts.map((orgAccount) => {
            return (
              <tr>
                <td>
                  <NavLink to={`/org-accounts/${orgAccount.id}`}>
                    {orgAccount.name}
                  </NavLink>{" "}
                  <sup>{renderCloudPlatform(orgAccount.cloudPlatform)}</sup>
                </td>
                <td>{orgAccount.cloudIdentifier}</td>
                <td>
                  <DeleteOrgUnitMembershipButton
                    orgUnitID={data?.orgUnit.id}
                    orgAccountID={orgAccount.id}
                    onCompleted={refetch}
                  />
                </td>
              </tr>
            );
          })}
        </tbody>
      </Table>
      <br />
      <NewOrgUnitMembershipButton
        orgDimension={data?.orgUnit.orgDimension}
        orgUnit={data?.orgUnit}
        orgAccount={undefined}
        key={data?.orgUnit.id}
        onCompleted={refetch}
      />
    </Container>
  );
};
