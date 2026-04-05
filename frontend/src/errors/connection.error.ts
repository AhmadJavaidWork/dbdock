import { DatabaseDriver } from "@/types/databaseDriver.types";
import { getErrorMessage } from "@/utils/errorMessagehelper";

export function connectionErrorToText(
  error: unknown,
  type: DatabaseDriver,
  username: string,
  databaseName: string
) {
  const message = getErrorMessage(error);
  switch (true) {
    case message.includes("no such host"):
      return "Host not found. Please check the hostname.";

    case message.includes("connection refused"):
      return "Connection refused. Check host and port.";

    case message.includes("password authentication failed"):
      return `Password authentication failed for user ${username}. Please check your password.`;

    case message.includes(`database "${databaseName}" does not exist`):
      return `Database "${databaseName}" does not exist`;

    default:
      return message;
  }
}
