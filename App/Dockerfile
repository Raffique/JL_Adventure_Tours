# Build stage
FROM node:21.7.3-alpine AS build-stage

WORKDIR /app

# Copy package.json and package-lock.json to install dependencies
COPY package.json package-lock.json ./
RUN npm install --no-optional

# Copy the rest of the Vue.js application files
COPY . .

#-------------------- UNCOMMMENT FOR DEVELOPMENT -------------------------
    # Expose the development server port
EXPOSE 3000

# Start the development server
CMD ["npm", "run", "dev"]

#-------------------- UNCOMMMENT FOR PRODUCTION -------------------------
# # Build the application for production
# RUN npm run build

# # Production stage
# FROM nginx:alpine AS production-stage

# # Copy the built files from the build stage to Nginx
# COPY --from=build-stage /app/dist /usr/share/nginx/html

# # Expose port 80 for the frontend
# EXPOSE 3000

# # Start Nginx
# CMD ["nginx", "-g", "daemon off;"]
