import React from "react";

import { title } from "@/components/primitives";
import DefaultLayout from "@/layouts/default";
import AreaTable from "@/pages/area/table-area";

export default function AreaPage() {
  return (
    <DefaultLayout>
      <section className="flex flex-col items-center justify-center h-full py-8 md:py-10">
        <div className="inline-block max-w-4xl w-full text-center justify-center">
          <h1 className={title({ color: "violet" })}>Gérez vos AREA</h1>
          <p className="text-large text-default-500 mt-2">
            gérez vos Action/Reaction intégrés a chaque services
          </p>
        </div>
        <div className="w-full max-w-4xl mt-8">
          <AreaTable />
        </div>
      </section>
    </DefaultLayout>
  );
}
