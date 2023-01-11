import React, { useState } from "react";
import { OrganizationalAccounts } from "../__generated__/graphql";
import { NavLink } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import { Table } from "react-bootstrap";
import Container from "react-bootstrap/Container";
import { renderCloudPlatformTableData } from "../utils/table_rendering";

const ORGANIZATIONAL_ACCOUNTS_QUERY = gql`
  query organizationalAccounts {
    organizationalAccounts(limit: 100) {
      items {
        orgAccountId
        name
        cloudPlatform
        cloudIdentifier
      }
    }
  }
`;

type Response = {
  organizationalAccounts: OrganizationalAccounts;
};

export const OrganizationalAccountsList = () => {
  const [orgAccountId, setOrgAccountId] = useState("");

  const { loading, error, data } = useQuery<Response>(
    ORGANIZATIONAL_ACCOUNTS_QUERY,
    {}
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container>
      <h1>Organizational Accounts</h1>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Org Account Id</th>
            <th>Org Account Name</th>
            <th>Cloud Platform</th>
            <th>Cloud ID</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalAccounts.items.map((account) => {
            return (
              <tr>
                <td>
                  <NavLink to={`/org-accounts/${account?.orgAccountId}`}>
                    {account?.orgAccountId}
                  </NavLink>
                </td>
                <td>{account?.name}</td>
                {renderCloudPlatformTableData(account?.cloudPlatform)}
                <td>{account?.cloudIdentifier}</td>
              </tr>
            );
          })}
        </tbody>
      </Table>
    </Container>
  );
};
