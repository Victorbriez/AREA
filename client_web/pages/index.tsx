import React from "react";
import { Link } from "@nextui-org/link";
import { button as buttonStyles } from "@nextui-org/theme";

import { title, subtitle } from "@/components/primitives";
import DefaultLayout from "@/layouts/default";

export default function IndexPage() {
  return (
    <DefaultLayout>
      <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10 px-4 sm:px-6 lg:px-8">
        <div className="inline-block max-w-xl text-center justify-center">
          <h1 className="text-3xl sm:text-4xl md:text-5xl font-bold mb-4">
            <span className={title()}>Bienvenue sur&nbsp;</span>
            <span className={title({ color: "violet" })}>AREA&nbsp;</span>
          </h1>
          <h2 className="text-2xl sm:text-3xl md:text-4xl font-semibold mb-4">
            <span className={title()}>
              Connectez vos services favoris et automatisez vos tâches.
            </span>
          </h2>
          <div className={subtitle({ class: "mt-4 text-lg sm:text-xl" })}>
            La plateforme d&apos;automatisation puissante et flexible.
          </div>
        </div>

        <div className="flex flex-col sm:flex-row gap-3 mt-6 w-full sm:w-auto">
          <Link
            className={buttonStyles({
              color: "secondary",
              radius: "full",
              variant: "shadow",
              class: "w-full sm:w-auto",
            })}
            href="/services"
          >
            Explorer les Services
          </Link>
          <Link
            className={buttonStyles({
              variant: "bordered",
              radius: "full",
              class: "w-full sm:w-auto",
            })}
            href="/area"
          >
            Créer une Nouvelle Automatisation
          </Link>
        </div>
      </section>
    </DefaultLayout>
  );
}
