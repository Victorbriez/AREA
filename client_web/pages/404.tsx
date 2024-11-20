import { Button } from "@nextui-org/react";
import { useRouter } from "next/router";

import { title, subtitle } from "@/components/primitives";

export default function Custom404() {
  const router = useRouter();

  return (
    <section className="flex items-center justify-center min-h-screen">
      <div className="flex flex-col items-center justify-center space-y-4">
        <h1 className={title({ color: "violet" })}>404</h1>
        <div className={subtitle({ class: "text-center" })}>
          La page que vous cherchez n&apos;existe pas.
        </div>
        <Button
          className="!mt-4"
          color="secondary"
          radius="full"
          variant="shadow"
          onClick={() => router.push("/")}
        >
          Retour à l&apos;accueil
        </Button>
      </div>
    </section>
  );
}
