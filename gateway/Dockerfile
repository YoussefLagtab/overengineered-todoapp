# base -------------------------------
FROM node:18-alpine as base

WORKDIR /app

# build -------------------------------
FROM base as build

COPY package.json yarn.lock .
RUN yarn

COPY . .
RUN yarn build
RUN yarn --production --ignore-scripts --prefer-offline --frozen-lockfile

# gateway -------------------------------
FROM base

COPY --from=build /app/package.json ./
COPY --from=build /app/node_modules/ ./node_modules
COPY --from=build /app/dist/ ./dist

CMD ["yarn", "start:prod"]
