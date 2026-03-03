from __future__ import annotations

import logfire
import typer

import log
from config import get_settings

app = typer.Typer()


@app.command()
def serve() -> None:
    """Start the application server."""
    settings = get_settings()
    log.setup(settings.app, settings.otel)

    logfire.info(
        "starting {name}",
        name=settings.app.name,
        config=settings.model_dump(),
    )

    # TODO: initialize and start your application here


if __name__ == "__main__":
    app()
