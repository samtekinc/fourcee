import React, { ReactNode } from "react";
import { NotificationContainer } from "react-notifications";
import { Header } from "./Header";
import "react-notifications/lib/notifications.css";
type BasicProps = {
  children: ReactNode[] | ReactNode;
};

export const PageWrapper = ({ children }: BasicProps) => {
  return (
    <article className="height-full">
      <Header />
      <NotificationContainer />
      <main className="app-container height-full">
        <div className="main-page">{children}</div>
      </main>
    </article>
  );
};
