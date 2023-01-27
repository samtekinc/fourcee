import { OrgDimension } from "../__generated__/graphql";
import { NavLink, Outlet, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Container from "react-bootstrap/Container";
import { Tree } from "react-org-chart";
import { NewOrgUnitButton } from "./NewOrgUnitButton";
import { GetOrgUnitTree, OrgUnitTreeNode } from "../utils/org_tree_rendering";
import { Breadcrumb } from "react-bootstrap";

const ORG_DIMENSION_QUERY = gql`
  query orgDimension($orgDimensionID: ID!) {
    orgDimension(orgDimensionID: $orgDimensionID) {
      id
      name
      rootOrgUnitID
      orgUnits {
        id
        name
        parentOrgUnitID
        hierarchy
      }
    }
  }
`;

type Response = {
  orgDimension: OrgDimension;
};

export const OrgDimensionPage = () => {
  const params = useParams();

  const orgDimensionID = params.orgDimensionID ? params.orgDimensionID : "";

  const { loading, error, data, refetch } = useQuery<Response>(
    ORG_DIMENSION_QUERY,
    {
      variables: {
        orgDimensionID: orgDimensionID,
      },
      pollInterval: 5000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  let orgUnitsMap = GetOrgUnitTree(
    orgDimensionID,
    data?.orgDimension.orgUnits ?? []
  );

  return (
    <Container
      style={{
        paddingTop: "2rem",
        paddingBottom: "2rem",
      }}
      fluid
    >
      <Breadcrumb>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/" }}>
          Home
        </Breadcrumb.Item>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/org-dimensions" }}>
          Organizations
        </Breadcrumb.Item>
        <Breadcrumb.Item active>
          {data?.orgDimension.name} ({data?.orgDimension.id})
        </Breadcrumb.Item>
      </Breadcrumb>
      <Container style={{ minWidth: "50vw" }}>
        <NewOrgUnitButton
          key={data?.orgDimension.id ?? ""}
          orgDimensionID={data?.orgDimension.id ?? ""}
          existingOrgUnits={data?.orgDimension.orgUnits ?? []}
          onCompleted={refetch}
        />
        <Tree
          label={<h1>{data?.orgDimension.name} Org Dimension</h1>}
          lineWidth={"2px"}
          nodePadding={"30px"}
        >
          <OrgUnitTreeNode
            orgUnit={orgUnitsMap.get(data?.orgDimension.rootOrgUnitID ?? "")}
          />
        </Tree>
        <br />
      </Container>
      <Outlet />
    </Container>
  );
};
