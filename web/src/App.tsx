import React from "react";
import { Routes, Route } from "react-router-dom";

import { PageWrapper } from "./components/PageWrapper";
import { OrgDimensionsList } from "./components/OrgDimensionsList";
import { OrgDimensionPage } from "./components/OrgDimensionPage";

import { OrgAccountsList } from "./components/OrgAccountsList";
import { OrgAccountPage } from "./components/OrgAccountPage";
import { OrgUnitPage } from "./components/OrgUnitPage";

import { ModuleGroupsList } from "./components/ModuleGroupList";
import { ModuleGroupPage } from "./components/ModuleGroupPage";
import { ModuleVersionPage } from "./components/ModuleVersionPage";

import { ModuleAssignmentsList } from "./components/ModuleAssignmentsList";
import { ModuleAssignmentPage } from "./components/ModuleAssignmentPage";

import { PlanExecutionRequestPage } from "./components/PlanExecutionRequestPage";
import { ApplyExecutionRequestPage } from "./components/ApplyExecutionRequestPage";

import { ModulePropagationExecutionRequestPage } from "./components/ModulePropagationExecutionPage";
import { ModulePropagationPage } from "./components/ModulePropagationPage";
import { ModulePropagationDriftCheckRequestPage } from "./components/ModulePropagationDriftCheckPage";
import { ModulePropagationsList } from "./components/ModulePropagationsList";

function App() {
  return (
    <PageWrapper>
      <Routes>
        <Route path="/" element={<OrgDimensionsList />} />
        <Route path="/module-assignments" element={<ModuleAssignmentsList />}>
          <Route
            path="/module-assignments/:moduleAssignmentID"
            element={<ModuleAssignmentPage />}
          />
        </Route>
        <Route
          path="/plan-execution-requests/:planExecutionRequestID"
          element={<PlanExecutionRequestPage />}
        />
        <Route
          path="/apply-execution-requests/:applyExecutionRequestID"
          element={<ApplyExecutionRequestPage />}
        />
        <Route path="/org-structures" element={<OrgDimensionsList />}>
          <Route
            path="/org-structures/:orgDimensionID"
            element={<OrgDimensionPage />}
          >
            <Route
              path="/org-structures/:orgDimensionID/org-units/:orgUnitID"
              element={<OrgUnitPage />}
            />
          </Route>
        </Route>
        <Route path="/org-accounts" element={<OrgAccountsList />}>
          <Route
            path="/org-accounts/:orgAccountID"
            element={<OrgAccountPage />}
          />
        </Route>
        <Route path="/module-groups" element={<ModuleGroupsList />}>
          <Route
            path="/module-groups/:moduleGroupID"
            element={<ModuleGroupPage />}
          />
          <Route
            path="/module-groups/:moduleGroupID/versions/:moduleVersionID"
            element={<ModuleVersionPage />}
          />
        </Route>
        <Route path="/module-propagations" element={<ModulePropagationsList />}>
          <Route
            path="/module-propagations/:modulePropagationID"
            element={<ModulePropagationPage />}
          >
            <Route
              path="/module-propagations/:modulePropagationID/executions/:modulePropagationExecutionRequestID"
              element={<ModulePropagationExecutionRequestPage />}
            />
            <Route
              path="/module-propagations/:modulePropagationID/drift-checks/:modulePropagationDriftCheckRequestID"
              element={<ModulePropagationDriftCheckRequestPage />}
            />
          </Route>
        </Route>
      </Routes>
    </PageWrapper>
  );
}

export default App;
