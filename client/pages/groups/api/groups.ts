import { Group } from "../../../models/group";
import { postHeader } from "../../../api/post";

export type getProps = {
	authorization: string;
};
export const clientHost = "http://localhost:3000/";

export const getHeader = (authorization: string) => {
	return new Headers({
		"Access-Control-Request-Method": "GET",
		"Access-Control-Request-Headers":
			"origin, content-type, accept, access-control-request-method, authorization",
		Origin: clientHost,
		"Content-Type": "application/json",
		Accept: "application/json",
		Authorization: authorization,
	});
};

export const fetchGroups = async (authentication: string) => {
	const data: RequestInit = {
		method: "GET",
		mode: "cors",
		headers: getHeader(authentication),
	};
	const response: Response = await fetch(
		process.env.NEXT_PUBLIC_BACKEND_HOST + "/groups",
		data
	);
	const responseBody = await response.json();

	if (response.status == 401) {
		if (responseBody.message === "Token is expired") {
			console.log("reflesh");
			// TODO:リフレッシュ処理に変更
			localStorage.clear()

		} else {
			console.log("redirect");
			localStorage.clear()
			// トークン破棄 & 認証リダイレクト
		}
	}

	return responseBody as { groups: Group[] };
};

type RegisterGroupProps = {
	authorization: string;
	groupId: string;
};

type RegisterGroupResponse = {
	group?: Group;
};

export const registerGroup = async ({
	authorization,
	groupId,
}: RegisterGroupProps) => {
	const body = {
		group_id: groupId,
	};
	const data: RequestInit = {
		method: "POST",
		mode: "cors",
		headers: { ...postHeader(), authorization },
		body: JSON.stringify(body),
	};

	console.log(process.env.NEXT_PUBLIC_BACKEND_HOST)
	const response = await fetch(
		process.env.NEXT_PUBLIC_BACKEND_HOST + "/groups/register",
		data
	);
	const responseBody = await response.json();

	if (response.status === 401) {
		if (responseBody.message === "Token is expired") {
			console.log("reflesh");
			// TODO:リフレッシュ処理に変更
			localStorage.clear()

		} else {
			console.log("redirect");
			localStorage.clear()

			// トークン破棄 & 認証リダイレクト
		}
	}

	if (isRegisterGroupResponse(responseBody)) {
		return responseBody;
	} else {
		const emptyResponse: RegisterGroupResponse = { group: undefined };
		return emptyResponse;
	}
};

const isRegisterGroupResponse = (
	response: any
): response is RegisterGroupResponse => {
	if (typeof response === "object") {
		return !!response.group;
	}

	return false;
};
