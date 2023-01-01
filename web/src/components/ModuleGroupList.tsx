import React, { useState } from "react";
import { ModuleGroups } from "../__generated__/graphql";
import { NavLink } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import { Table } from "react-bootstrap";
import Container from "react-bootstrap/Container";

const MODULE_GROUPS_QUERY = gql`
  query moduleGroups {
    moduleGroups(limit: 100) {
      items {
        moduleGroupId
        name
      }
    }
  }
`;

type Response = {
  moduleGroups: ModuleGroups;
};

export const ModuleGroupsList = () => {
  const [orgAccountId, setOrgAccountId] = useState("");

  const { loading, error, data } = useQuery<Response>(MODULE_GROUPS_QUERY, {});

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container>
      <h1>Module Groups</h1>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Module Group Name</th>
            <th>Module Group ID</th>
          </tr>
        </thead>
        <tbody>
          {data?.moduleGroups.items.map((moduleGroup) => {
            return (
              <tr>
                <td>{moduleGroup?.name}</td>
                <td>
                  <NavLink to={`/module-groups/${moduleGroup?.moduleGroupId}`}>
                    {moduleGroup?.moduleGroupId}
                  </NavLink>
                </td>
              </tr>
            );
          })}
        </tbody>
      </Table>
    </Container>
  );
};
